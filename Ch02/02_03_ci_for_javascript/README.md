# 02_03 CI for JavaScript

Check node versions in the matrix strategy.

Update the matrix according to the versions needed for your project and the active versions as reported by [Node.js Release Working Group](https://github.com/nodejs/Release).

For example, change:

    strategy:
      matrix:
        node-version: [14.x, 16.x, 18.x]

to:


    strategy:
      matrix:
        node-version: [16.x, 18.x, 19.x]


