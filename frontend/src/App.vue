
<template>
  <v-app>
    <v-layout>
      <v-container class="container">
        <Header></Header>
        <Display></Display>
        <Footer></Footer>
      </v-container>
    </v-layout>
  </v-app>
</template>

<script setup lang="ts">
import Header from "./components/header.vue"
import Display from "./components/display.vue"
import Footer from "./components/footer.vue"
</script>

<style scoped>
.container {
  height: 100vh;
  display: flex;
  flex-direction: column;
}
</style>
