<template>
  <div class="hello">
    <div>{{this.$store.state.text}}</div>
    <div>{{this.$store.state.num}}</div>
    <button v-on:click="pushedButton(10)">up</button>
  </div>
</template>

<script>
export default {
  methods: {
    pushedButton(num) {
      this.$store.dispatch("addNum", { num });
    }
  }
};
</script>

<style scoped>
</style>
