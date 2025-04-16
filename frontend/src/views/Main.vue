<template>
    <div>
      <main-top></main-top>
      <main-bottom></main-bottom>
    </div>
</template>
<script>
import MainTop from "./components/MainTop.vue";
import MainBottom from "./components/MainBottom.vue";

export default {
  name: "components",
  components: {
    MainTop,
    MainBottom,
  }
};
</script>
