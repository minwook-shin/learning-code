<template>
  <div>
    <input type="text" v-model="Item" v-on:keyup.enter="addItem">
    <button v-on:click="addItem">
      <i class="fas fa-plus"></i>
    </button>
  </div>
</template>

<script>
export default {
  data: function() {
    return {
      Item: ""
    };
  },
  methods: {
    addItem: function() {
      this.$emit("AddTodoItem", this.Item);
      this.clearItem();
    },
    clearItem: function() {
      this.Item = "";
    }
  }
};
</script>

<style scoped>
</style>
