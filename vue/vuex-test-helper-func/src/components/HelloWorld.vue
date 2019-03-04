<template>
  <div class="hello">
    <div>{{text}}</div>
    <div>{{num}}</div>
    <button v-on:click="pushedButton(10)">up</button>
    <button v-on:click="pushedButton2(10)">up</button>
  </div>
</template>

<script>
import { mapActions } from "vuex";
import { mapGetters } from "vuex";
import { mapMutations } from "vuex";
import { mapState } from "vuex";

export default {
  methods: {
    ...mapActions(["addNum"]),
    pushedButton(num) {
      this.addNum({
        num: num
      });
    },
    ...mapMutations({ add: "addNum" }),
    pushedButton2(num) {
      this.add({
        value: num
      });
    }
  },
  computed: {
    ...mapGetters(["text"]),
    ...mapState(["num"])
  }
};
</script>

<style scoped>
</style>
