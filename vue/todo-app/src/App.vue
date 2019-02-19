<template>
  <div id="app">
    <todo-header></todo-header>
    <todo-input v-on:AddTodoItem="addOneItem"></todo-input>
    <todo-list v-bind:propsData="todoItems" v-on:RemoveTodoItem="removeOneItem"></todo-list>
    <todo-footer v-on:ClearAll="clearAllItem"></todo-footer>
  </div>
</template>

<script>
import TodoHeader from "./components/TodoHeader";
import TodoInput from "./components/TodoInput";
import TodoList from "./components/TodoList";
import TodoFooter from "./components/TodoFooter";

export default {
  data() {
    return {
      todoItems: []
    };
  },
  methods: {
    addOneItem: function(i) {
      localStorage.setItem(i, i);
      this.todoItems.push(i);
    },
    removeOneItem: function(item, index) {
      localStorage.removeItem(item);
      this.todoItems.splice(index, 1);
    },
    clearAllItem: function() {
      localStorage.clear();
      this.todoItems = [];
    }
  },
  components: {
    TodoHeader,
    TodoInput,
    TodoList,
    TodoFooter
  },
  created() {
    if (localStorage.length > 0) {
      for (var i = 0; i < localStorage.length; i++) {
        if (localStorage.key(i) != "loglevel:webpack-dev-server")
          this.todoItems.push(localStorage.key(i));
      }
    }
  }
};
</script>

<style>
body {
  text-align: center;
  background-color: mintcream;
}

input {
  border-style: groove;
  width: 60%;
}

button {
  border-style: groove;
  padding: 5px;
  margin: 10px;
}
.shawdow {
  box-shadow: 5px 10px 5px;
}
</style>
