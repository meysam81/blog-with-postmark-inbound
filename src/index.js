import "@/styles.css";
import log from "loglevel";

var publishBtn;
var publishInfoPanel;
var closePublishInfoBtn;
var copyEmailBtn;
var emailToCopy;

function formatISODate(dateString) {
  return new Date(dateString).toISOString();
}

function formatReadableDate(dateString) {
  return new Date(dateString).toLocaleDateString("en-US", {
    year: "numeric",
    month: "long",
    day: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  });
}

function getAuthorInitial(email) {
  return email.charAt(0).toUpperCase();
}

async function fetchPosts() {
  try {
    const response = await fetch("/api/posts");
    if (!response.ok) {
      throw new Error("Failed to fetch posts");
    }
    return await response.json();
  } catch (error) {
    log.error("Error fetching posts:", error);
    return [];
  }
}

function renderPosts(posts) {
  const mainContent = document.getElementById("main-content");

  if (!posts || posts?.length == 0) {
    return;
  }

  mainContent.innerHTML = `
    <div class="max-w-4xl mx-auto">
      <div class="grid gap-8">
        ${posts
          .map(
            (post) => `
          <article class="blog-card bg-white dark:bg-gray-800 rounded-xl shadow-md p-8 fade-in border-l-4 border-primary-500">
            <h2 class="text-2xl md:text-3xl font-semibold text-gray-900 dark:text-gray-100 mb-3 font-display">
              ${post.title}
            </h2>
            <div class="prose prose-green dark:prose-invert max-w-none my-4">
              ${post.content}
            </div>
            <div class="flex items-center mt-6 pt-4 border-t border-gray-100 dark:border-gray-700">
              <div class="bg-primary-100 dark:bg-primary-900 text-primary-700 dark:text-primary-300 rounded-full h-10 w-10 flex items-center justify-center font-medium">
                ${getAuthorInitial(post["author-email"])}
              </div>
              <div class="ml-3">
                <p class="text-sm text-gray-700 dark:text-gray-300 font-medium">
                  ${post["author-email"]}
                </p>
                <p class="text-xs text-gray-500 dark:text-gray-400">
                  <time datetime="${formatISODate(
                    post["created-at"]
                  )}">${formatReadableDate(post["created-at"])}</time>
                </p>
              </div>
            </div>
          </article>
        `
          )
          .join("")}
      </div>
    </div>
  `;
}

function wrapCharsWithSpan() {
  var titleElement = document.getElementById("titleText");
  var titleText = titleElement.textContent.trim();
  var wrappedTitle = "";

  for (var i = 0; i < titleText.length; i++) {
    var char = titleText[i];
    if (char === " ") {
      wrappedTitle += " ";
    } else {
      wrappedTitle +=
        '<span class="char-hover" aria-hidden="true">' + char + "</span>";
    }
  }

  titleElement.innerHTML = wrappedTitle;

  var srText = document.createElement("span");
  srText.className = "sr-only";
  srText.textContent = titleText;
  titleElement.appendChild(srText);
}

function setupYearInFooter() {
  var yearElement = document.getElementById("current-year");
  yearElement.textContent = new Date().getFullYear();
}

function setupPublishInfo() {
  publishBtn = document.getElementById("publishBtn");
  publishInfoPanel = document.getElementById("publishInfoPanel");
  closePublishInfoBtn = document.getElementById("closePublishInfoBtn");
  copyEmailBtn = document.getElementById("copyEmailBtn");
  emailToCopy = "blog@tarzan.meysam.io";
}

function togglePublishInfo() {
  var isExpanded = publishInfoPanel.classList.contains("active");
  publishInfoPanel.classList.toggle("active");
  publishBtn.setAttribute("aria-expanded", !isExpanded);
  publishInfoPanel.setAttribute("aria-hidden", isExpanded);
}

function closePublishInfo() {
  publishInfoPanel.classList.remove("active");
  publishBtn.setAttribute("aria-expanded", "false");
  publishInfoPanel.setAttribute("aria-hidden", "true");
}

function copyEmailToClipboard() {
  navigator.clipboard.writeText(emailToCopy).then(function () {
    var originalTitle = copyEmailBtn.getAttribute("title");
    copyEmailBtn.setAttribute("title", "Copied!");

    copyEmailBtn.innerHTML =
      '<svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-primary-500" viewBox="0 0 20 20" fill="currentColor"><path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd" /></svg>';

    setTimeout(function () {
      copyEmailBtn.setAttribute("title", originalTitle);
      copyEmailBtn.innerHTML =
        '<svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor"><path d="M8 3a1 1 0 011-1h2a1 1 0 110 2H9a1 1 0 01-1-1z" /><path d="M6 3a2 2 0 00-2 2v11a2 2 0 002 2h8a2 2 0 002-2V5a2 2 0 00-2-2 3 3 0 01-3 3H9a3 3 0 01-3-3z" /></svg>';
    }, 2000);

    var announcement = document.createElement("div");
    announcement.setAttribute("aria-live", "polite");
    announcement.className = "sr-only";
    announcement.textContent = "Email address copied to clipboard";
    document.body.appendChild(announcement);

    setTimeout(function () {
      document.body.removeChild(announcement);
    }, 3000);
  });
}

function main() {
  log.debug("Starting the app...");

  setupPublishInfo();

  publishBtn.addEventListener("click", togglePublishInfo);
  closePublishInfoBtn.addEventListener("click", closePublishInfo);
  copyEmailBtn.addEventListener("click", copyEmailToClipboard);

  document.addEventListener("click", function (event) {
    var isClickInside =
      publishInfoPanel.contains(event.target) ||
      publishBtn.contains(event.target);
    if (!isClickInside && publishInfoPanel.classList.contains("active")) {
      closePublishInfo();
    }
  });

  document.addEventListener("keydown", function (event) {
    if (
      event.key === "Escape" &&
      publishInfoPanel.classList.contains("active")
    ) {
      closePublishInfo();
    }
  });

  wrapCharsWithSpan();
  setupPublishInfo();
  setupYearInFooter();

  fetchPosts().then((posts) => {
    renderPosts(posts);
  });
}

if (import.meta.env.DEV) {
  log.setDefaultLevel("debug");
} else {
  log.setDefaultLevel("warn");
}

log.info("Tarzan is running!");
document.addEventListener("DOMContentLoaded", main);
