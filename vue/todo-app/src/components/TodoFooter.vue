<template>
  <div>
    <button v-on:click="clearList">Clear All</button>
  </div>
</template>

<script>
export default {
  methods: {
    clearList: function() {
      this.$emit("ClearAll");
    }
  }
};
</script>

<style>
</style>
