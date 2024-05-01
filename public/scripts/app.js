function resize() {
  const el = document.getElementById("autoresize")
  if (!el) {
    return
  }
  el.style.height = "auto"
  el.style.height = el.scrollHeight + "px"
}

// When the DOM has finished loading, but not images etc.
document.addEventListener('DOMContentLoaded', resize)

// Input event on input, select, textarea elements.
document.addEventListener('input', (event) => {
  if (event.target.id === "autoresize") {
    resize()
  }
})

// Make it work with HTMX as well.
document.body.addEventListener('htmx:afterSwap', (event) => {
  if (event.detail.target === document.body) {
    resize()
  }
})

window.addEventListener("resize", resize)
