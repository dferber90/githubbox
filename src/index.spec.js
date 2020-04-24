import test from 'tape'
import { getCodeSandboxLocation } from './location.js'

test('root', t => {
  t.equals(getCodeSandboxLocation('/'), null)
  t.end()
})

test('repo root', t => {
  t.equals(
    getCodeSandboxLocation('/zeit/ms'),
    'https://codesandbox.io/s/github/zeit/ms',
  )
  t.end()
})

test('repo branch', t => {
  t.equals(
    getCodeSandboxLocation('/zeit/ms/tree/master'),
    'https://codesandbox.io/s/github/zeit/ms/tree/master',
  )
  t.end()
})

test('repo tree', t => {
  t.equals(
    getCodeSandboxLocation('/zeit/ms/tree/2.1.1'),
    'https://codesandbox.io/s/github/zeit/ms/tree/2.1.1',
  )
  t.end()
})

test('repo commit', t => {
  t.equals(
    getCodeSandboxLocation(
      '/zeit/ms/tree/7920885eb232fbe7a5efdab956d3e7c507c92ddf',
    ),
    'https://codesandbox.io/s/github/zeit/ms/tree/7920885eb232fbe7a5efdab956d3e7c507c92ddf',
  )
  t.end()
})

test('branch with file', t => {
  t.equals(
    getCodeSandboxLocation('/zeit/ms/blob/master/package.json'),
    'https://codesandbox.io/s/github/zeit/ms/tree/master?file=/package.json',
  )
  t.end()
})

test('tree with file', t => {
  t.equals(
    getCodeSandboxLocation('/zeit/ms/blob/2.0.0/index.js'),
    'https://codesandbox.io/s/github/zeit/ms/tree/2.0.0?file=/index.js',
  )
  t.end()
})

test('commit with file', t => {
  t.equals(
    getCodeSandboxLocation(
      '/zeit/ms/blob/adf1eb282d29fe3c405d205a3854177b86a97c1f/index.js',
    ),
    'https://codesandbox.io/s/github/zeit/ms/tree/adf1eb282d29fe3c405d205a3854177b86a97c1f?file=/index.js',
  )
  t.end()
})

test('pages which are not repos', t => {
  t.equals(getCodeSandboxLocation('/issues'), null)
  t.equals(getCodeSandboxLocation('/zeit/ms/stargazers'), null)
  t.end()
})
