import { getCodeSandboxLocation } from './location.js'

async function handleEvent(event) {
  const url = new URL(event.request.url)

  let path = url.pathname

  if (path === '/' || path === '/index.html')
    return Response.redirect('https://github.com/dferber90/githubbox', 302)

  if (path === '/robots.txt')
    return new Response(
      [
        '# https://www.robotstxt.org/robotstxt.html',
        'User-agent: *',
        'Disallow:',
      ].join('\n'),
      { status: 200 },
    )

  let location = getCodeSandboxLocation(path)

  if (location) {
    return Response.redirect(location, 302)
  }

  return new Response('Not found', { status: 404 })
}

export function handleFetch(event) {
  try {
    event.respondWith(handleEvent(event))
  } catch (e) {
    event.respondWith(new Response('Internal Error', { status: 500 }))
  }
}

// disable during testing
if (typeof addEventListener === 'function')
  addEventListener('fetch', handleFetch)
