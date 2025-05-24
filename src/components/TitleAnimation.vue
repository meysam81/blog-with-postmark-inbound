<template>
  <div
    class="text-center text-gray-700 dark:text-gray-200 mt-4 text-2xl font-medium px-4"
    ref="titleElement"
  >
    {{ titleText }}
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'

export default {
  name: 'TitleAnimation',
  setup() {
    var titleElement = ref(null)
    var titleText = "Transforming Your Emails into Blogs!"

    function wrapCharsWithSpan() {
      if (!titleElement.value) {return}

      var wrappedTitle = ""

      for (var i = 0; i < titleText.length; i++) {
        var char = titleText[i]
        if (char === " ") {
          wrappedTitle += '<span class="char-hover"> </span>'
        } else {
          wrappedTitle += `<span class="char-hover">${char}</span>`
        }
      }

      titleElement.value.innerHTML = wrappedTitle

      var srText = document.createElement("span")
      srText.className = "sr-only"
      srText.textContent = titleText
      titleElement.value.appendChild(srText)
    }

    onMounted(function() {
      wrapCharsWithSpan()
    })

    return {
      titleElement,
      titleText
    }
  }
}
</script>
