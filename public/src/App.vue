<template>
  <div id="app">
    <navigator></navigator>
    <router-view></router-view>
    <!--<img alt="Vue logo" src="./assets/logo.png">-->
    <!--<HelloWorld msg="Welcome to Your Vue.js App"/>-->
    <!--<file-index></file-index>-->
    <!--<login></login>-->
  </div>
</template>

<script>
// import HelloWorld from './components/HelloWorld.vue'
import Navigator from "./components/Nav";
// import Login from './components/Login'
// import FileIndex from './components/files/ListFiles'

export default {
  name: 'app',
  components: {
    Navigator,
    // FileIndex,
    // Login,
  }
}
</script>

<style>
/*#app {*/
  /*font-family: 'Avenir', Helvetica, Arial, sans-serif;*/
  /*!*-webkit-font-smoothing: antialiased;*!*/
  /*!*-moz-osx-font-smoothing: grayscale;*!*/
  /*text-align: center;*/
  /*!*color: #2c3e50;*!*/
/*}*/
</style>
