<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
html,
body,
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  height: 100%;
}

html,
body {
  margin: 0px;
  padding: 0px;
  height: 100%;
  width: 100%;
}

.app-container {
  height: 100%;
}

.main-part {
  padding: 0;
}
</style>
