# envexp
Export environment variables to a json or javascript file

## Use case
One use case for this small cli tool is to enable runtime variables when deploying a create-react-app web-app with Docker. The idea consists in using a `env.js` script included from `index.html`. This script sets the `window._env` property with environment variables. However, this script must be created during start up of your docker container.

With `envexp` tool you can create the `env.js` script easily before start serving your web-app.

Fore more details read this [issue](https://github.com/facebook/create-react-app/issues/2353).

## Usage Examples:

### Basic export
```
./envexp -t json > env.json
./envexp -t web > env.js
```
### Export environment variables starting with prefix
```
./envexp -t json -prefix REACT_APP_ > env.json
./envexp -t web -prefix REACT_APP_ > env.js
```

### Basic export with pretty output
```
./envexp -p -t json > env.json
./envexp -p -t web > env.js
```