import log from 'loglevel'

function formatISODate(dateString) {
  return new Date(dateString).toISOString()
}

function formatReadableDate(dateString) {
  return new Date(dateString).toLocaleDateString("en-US", {
    year: "numeric",
    month: "long",
    day: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  })
}

function getAuthorInitial(email) {
  return email.charAt(0).toUpperCase()
}

async function fetchPosts() {
  try {
    var response = await fetch("/api/posts")
    if (!response.ok) {
      throw new Error("Failed to fetch posts")
    }
    return await response.json()
  } catch (error) {
    log.error("Error fetching posts:", error)
    return []
  }
}

export {
  formatISODate,
  formatReadableDate,
  getAuthorInitial,
  fetchPosts
}
