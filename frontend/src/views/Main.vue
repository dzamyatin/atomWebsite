<template>
    <div>
<!--        <hero></hero>-->
      <main-top></main-top>
      <main-bottom></main-bottom>
    </div>
</template>
<script>
import MainTop from "./components/MainTop.vue";
import MainBottom from "./components/MainBottom.vue";
// import Hero from "./components/Hero";

export default {
  name: "components",
  components: {
    MainTop,
    MainBottom,
    // Hero,
  }
};
</script>
