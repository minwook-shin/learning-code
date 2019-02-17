<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png">
    <NewLocalWorld></NewLocalWorld>
    <NewGlobalWorld></NewGlobalWorld>
  </div>
</template>

<script>
import NewLocalWorld from './components/NewLocalWorld';

export default {
  name: 'app',
  components: {
    NewLocalWorld
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
