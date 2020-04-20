import { getAssetFromKV } from '@cloudflare/kv-asset-handler'


function getCodeSandboxLocation(path) {
  const prefix = 'https://codesandbox.io/s/github/'
  const parts = path.substr(1).split('/')
  
  if (parts.length === 0) return null
  if (parts.length === 1) return null
  if (parts.length === 2) return prefix + parts.join('/')
  if (parts.length === 3) return null

  if (parts[2] === 'tree') return prefix + parts.join('/')

  if (parts[2] === 'blob') return prefix + [
    parts[0],
    parts[1],
    'tree',
    parts[3]
  ].join('/') + '?file=/' + parts.slice(4).join('/')

  return null
}

/**
 * The DEBUG flag will do two things that help during development:
 * 1. we will skip caching on the edge, which makes it easier to
 *    debug.
 * 2. we will return an error message on exception in your Response rather
 *    than the default 404.html page.
 */
const DEBUG = true

addEventListener('fetch', event => {
  try {
    event.respondWith(handleEvent(event))
  } catch (e) {
    if (DEBUG) {
      return event.respondWith(
        new Response(e.message || e.toString(), {
          status: 500,
        }),
      )
    }
    event.respondWith(new Response('Internal Error', { status: 500 }))
  }
})

async function handleEvent(event) {
  const url = new URL(event.request.url)
  let options = {}

  /**
   * You can add custom logic to how we fetch your assets
   * by configuring the function `mapRequestToAsset`
   */
  // options.mapRequestToAsset = handlePrefix(/^\/docs/)

  try {
    if (DEBUG) {
      // customize caching
      options.cacheControl = {
        bypassCache: true,
      }
    }
    return await getAssetFromKV(event, options)
  } catch (e) {
    let path = url.pathname
    let location = getCodeSandboxLocation(path)

    if (location) {
      return Response.redirect(location, 302 /* use 301 once its working */)
    }

    // redirect to landing page since wrong url was used
    return Response.redirect('https://githubbox.com/?404', 302)

    // if an error is thrown try to serve the asset at 404.html
    // if (!DEBUG) {
    //   try {
    //     let notFoundResponse = await getAssetFromKV(event, {
    //       mapRequestToAsset: req => new Request(`${new URL(req.url).origin}/404.html`, req),
    //     })

    //     return new Response(notFoundResponse.body, { ...notFoundResponse, status: 404 })
    //   } catch (e) {}
    // }

    // return new Response(e.message || e.toString(), { status: 500 })
  }
}
