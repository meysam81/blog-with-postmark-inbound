import { ref, onMounted, onUnmounted } from "vue";
import log from "loglevel";

const announcements = ref([]);
const focusedElement = ref(null);

function announceToScreenReader(message, priority = "polite") {
  if (!message || typeof message !== "string") {
    log.warn("Invalid message for screen reader announcement:", message);
    return;
  }

  var announcement = {
    id: Date.now(),
    message: message,
    priority: priority,
    timestamp: new Date(),
  };

  announcements.value.push(announcement);
  log.debug("Screen reader announcement:", announcement);

  // Remove announcement after a delay to prevent accumulation
  setTimeout(function removeAnnouncement() {
    announcements.value = announcements.value.filter(
      function filterAnnouncement(a) {
        return a.id !== announcement.id;
      }
    );
  }, 5000);
}

function focusElement(selector, options = {}) {
  try {
    var element =
      typeof selector === "string"
        ? document.querySelector(selector)
        : selector;

    if (!element) {
      log.warn("Element not found for focus:", selector);
      return false;
    }

    if (options.preventScroll) {
      element.focus({ preventScroll: true });
    } else {
      element.focus();
      if (options.scrollIntoView !== false) {
        element.scrollIntoView({
          behavior: "smooth",
          block: "center",
          inline: "nearest",
        });
      }
    }

    focusedElement.value = element;
    log.debug("Focused element:", element);
    return true;
  } catch (error) {
    log.error("Error focusing element:", error);
    return false;
  }
}

function trapFocus(containerSelector) {
  var container = document.querySelector(containerSelector);
  if (!container) {
    log.warn("Focus trap container not found:", containerSelector);
    return null;
  }

  var focusableSelectors = [
    "a[href]",
    "button:not([disabled])",
    "textarea:not([disabled])",
    "input:not([disabled])",
    "select:not([disabled])",
    '[tabindex]:not([tabindex="-1"])',
  ].join(", ");

  var focusableElements = container.querySelectorAll(focusableSelectors);
  var firstElement = focusableElements[0];
  var lastElement = focusableElements[focusableElements.length - 1];

  function handleTabKey(event) {
    if (event.key !== "Tab") {
      return;
    }

    if (event.shiftKey) {
      if (document.activeElement === firstElement) {
        event.preventDefault();
        lastElement.focus();
      }
    } else if (document.activeElement === lastElement) {
      event.preventDefault();
      firstElement.focus();
    }
  }

  function handleEscapeKey(event) {
    if (event.key === "Escape") {
      removeFocusTrap();
    }
  }

  function removeFocusTrap() {
    container.removeEventListener("keydown", handleTabKey);
    document.removeEventListener("keydown", handleEscapeKey);
    log.debug("Focus trap removed from:", containerSelector);
  }

  container.addEventListener("keydown", handleTabKey);
  document.addEventListener("keydown", handleEscapeKey);

  if (firstElement) {
    firstElement.focus();
  }

  log.debug("Focus trap enabled for:", containerSelector);
  return { remove: removeFocusTrap };
}

function skipToContent(targetSelector = "#main-content") {
  var target = document.querySelector(targetSelector);
  if (target) {
    target.focus();
    target.scrollIntoView({ behavior: "smooth", block: "start" });
    announceToScreenReader("Skipped to main content");
    return true;
  }
  log.warn("Skip target not found:", targetSelector);
  return false;
}

function enhanceKeyboardNavigation() {
  function handleGlobalKeyboard(event) {
    // Alt + S: Skip to content
    if (event.altKey && event.key === "s") {
      event.preventDefault();
      skipToContent();
      return;
    }

    // Alt + H: Focus on header/navigation
    if (event.altKey && event.key === "h") {
      event.preventDefault();
      focusElement('header, nav, [role="banner"]');
      announceToScreenReader("Focused on header navigation");
      return;
    }

    // Alt + M: Focus on main content
    if (event.altKey && event.key === "m") {
      event.preventDefault();
      focusElement('main, [role="main"]');
      announceToScreenReader("Focused on main content");
      return;
    }

    // Alt + F: Focus on footer
    if (event.altKey && event.key === "f") {
      event.preventDefault();
      focusElement('footer, [role="contentinfo"]');
      announceToScreenReader("Focused on footer");
      return;
    }
  }

  document.addEventListener("keydown", handleGlobalKeyboard);

  return function removeKeyboardNavigation() {
    document.removeEventListener("keydown", handleGlobalKeyboard);
  };
}

function checkColorContrast(foreground, background) {
  function getLuminance(color) {
    var rgb = color.match(/\d+/g);
    if (!rgb || rgb.length < 3) {
      return 0;
    }

    var r = parseInt(rgb[0], 10) / 255
    var g = parseInt(rgb[1], 10) / 255
    var b = parseInt(rgb[2], 10) / 255

    r = r <= 0.03928 ? r / 12.92 : ((r + 0.055) / 1.055) ** 2.4
    g = g <= 0.03928 ? g / 12.92 : ((g + 0.055) / 1.055) ** 2.4
    b = b <= 0.03928 ? b / 12.92 : ((b + 0.055) / 1.055) ** 2.4

    return 0.2126 * r + 0.7152 * g + 0.0722 * b;
  }

  var luminance1 = getLuminance(foreground);
  var luminance2 = getLuminance(background);
  var ratio =
    (Math.max(luminance1, luminance2) + 0.05) /
    (Math.min(luminance1, luminance2) + 0.05);

  return {
    ratio: ratio,
    aa: ratio >= 4.5,
    aaa: ratio >= 7,
  };
}

function useAccessibility() {
  var keyboardNavigationCleanup = null;

  onMounted(function onMountedHandler() {
    keyboardNavigationCleanup = enhanceKeyboardNavigation();

    // Add skip link if it doesn't exist
    var skipLink = document.querySelector(".skip-link");
    if (!skipLink) {
      skipLink = document.createElement("a");
      skipLink.href = "#main-content";
      skipLink.className = "skip-link";
      skipLink.textContent = "Skip to main content";
      skipLink.style.cssText = `
        position: absolute;
        top: -40px;
        left: 6px;
        background: #000;
        color: #fff;
        padding: 8px;
        z-index: 1000;
        text-decoration: none;
        border-radius: 4px;
        transition: top 0.3s;
      `;

      skipLink.addEventListener("focus", function onSkipLinkFocus() {
        skipLink.style.top = "6px";
      });

      skipLink.addEventListener("blur", function onSkipLinkBlur() {
        skipLink.style.top = "-40px";
      });

      skipLink.addEventListener("click", function onSkipLinkClick(event) {
        event.preventDefault();
        skipToContent();
      });

      document.body.insertBefore(skipLink, document.body.firstChild);
    }

    log.debug("Accessibility enhancements initialized");
  });

  onUnmounted(function onUnmountedHandler() {
    if (keyboardNavigationCleanup) {
      keyboardNavigationCleanup();
    }
  });

  return {
    announcements,
    focusedElement,
    announceToScreenReader,
    focusElement,
    trapFocus,
    skipToContent,
    checkColorContrast,
  };
}

export {
  useAccessibility,
  announceToScreenReader,
  focusElement,
  trapFocus,
  skipToContent,
  checkColorContrast,
};
