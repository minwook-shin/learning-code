<template>
  <div>
    <ul>
      <li v-bind:key="item" v-for="(item,index) in propsData">
        {{item}}
        <button v-on:click="removeItem(item, index)">
          <i class="fas fa-trash"></i>
        </button>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  methods: {
    removeItem: function(item, index) {
      this.$emit("RemoveTodoItem", item, index);
    }
  },
  props: ["propsData"]
};
</script>

<style>
</style>
