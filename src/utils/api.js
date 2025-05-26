import log from "loglevel";

function formatISODate(dateString) {
  try {
    if (!dateString) {
      return new Date().toISOString();
    }
    var date = new Date(dateString);
    if (isNaN(date.getTime())) {
      log.warn("Invalid date string provided:", dateString);
      return new Date().toISOString();
    }
    return date.toISOString();
  } catch (error) {
    log.error("Error formatting ISO date:", error);
    return new Date().toISOString();
  }
}

function formatReadableDate(dateString) {
  try {
    if (!dateString) {
      return "Unknown date";
    }
    var date = new Date(dateString);
    if (isNaN(date.getTime())) {
      log.warn("Invalid date string provided:", dateString);
      return "Unknown date";
    }
    return date.toLocaleDateString("en-US", {
      year: "numeric",
      month: "long",
      day: "numeric",
    });
  } catch (error) {
    log.error("Error formatting readable date:", error);
    return "Unknown date";
  }
}

function getAuthorInitial(authorName, fallbackEmail) {
  try {
    var nameToUse = authorName || fallbackEmail;
    if (!nameToUse || typeof nameToUse !== "string" || nameToUse.length === 0) {
      return "U"; // Default to "U" for Unknown
    }

    // If it's a name (contains space), get initials from first and last name
    if (nameToUse.includes(" ")) {
      var nameParts = nameToUse.trim().split(" ");
      if (nameParts.length >= 2) {
        return (nameParts[0].charAt(0) + nameParts[nameParts.length - 1].charAt(0)).toUpperCase();
      }
    }

    // For single word or email, just use first character
    return nameToUse.charAt(0).toUpperCase();
  } catch (error) {
    log.error("Error getting author initial:", error);
    return "U";
  }
}

async function fetchPosts() {
  try {
    log.debug("Fetching posts from /api/posts");
    var response = await fetch("/api/posts", {
      method: "GET",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      var errorMessage = `HTTP ${response.status}: ${response.statusText}`;
      log.error("Failed to fetch posts:", errorMessage);

      // Try to get error details from response
      try {
        var errorData = await response.text();
        if (errorData) {
          log.error("Server error details:", errorData);
        }
      } catch (parseError) {
        log.warn("Could not parse error response:", parseError);
      }

      throw new Error(`Failed to fetch posts: ${errorMessage}`);
    }

    var posts = await response.json();

    if (posts) {
      log.debug(`Successfully fetched ${posts.length} posts`);
    }

    // Validate posts structure
    if (!Array.isArray(posts)) {
      log.warn("Posts response is not an array, returning empty array");
      return [];
    }

    // Validate each post has required fields
    var validPosts = posts.filter(function validatePost(post) {
      if (!post || typeof post !== "object") {
        log.warn("Invalid post object found, skipping:", post);
        return false;
      }

      if (
        !post.title ||
        !post.content ||
        !post["author-email"] ||
        !post["created-at"]
      ) {
        log.warn("Post missing required fields, skipping:", post);
        return false;
      }

      return true;
    });

    if (validPosts.length !== posts.length) {
      log.warn(
        `Filtered out ${posts.length - validPosts.length} invalid posts`
      );
    }

    return validPosts;
  } catch (error) {
    if (error.name === "TypeError" && error.message.includes("fetch")) {
      log.error("Network error - server may be unavailable:", error.message);
    } else if (error.name === "SyntaxError") {
      log.error("Invalid JSON response from server:", error.message);
    } else {
      log.error("Error fetching posts:", error.message);
    }

    // Return empty array to prevent UI breakage
    return [];
  }
}

function fetchPostById(postId) {
  return fetch(`/api/posts/${postId}`, {
    method: "GET",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
  })
    .then(function handleResponse(response) {
      if (!response.ok) {
        var errorMessage = `HTTP ${response.status}: ${response.statusText}`;
        log.error("Failed to fetch post:", errorMessage);

        // Try to get error details from response
        try {
          return response.text().then(function parseErrorText(errorData) {
            if (errorData) {
              log.error("Server error details:", errorData);
            }
            throw new Error(`Failed to fetch post: ${errorMessage}`);
          });
        } catch (parseError) {
          log.warn("Could not parse error response:", parseError);
          throw new Error(`Failed to fetch post: ${errorMessage}`);
        }
      }
      return response.json()
    })
    .catch(function handleError(error) {
      log.error('Error fetching post:', error)
      throw error
    })
}

export { formatISODate, formatReadableDate, getAuthorInitial, fetchPosts, fetchPostById };
