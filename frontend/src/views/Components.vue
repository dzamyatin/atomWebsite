<template>
    <div>
        <hero></hero>
        <basic-elements></basic-elements>
        <inputs></inputs>
        <custom-controls></custom-controls>
        <navigation></navigation>
        <javascript-components></javascript-components>
        <icons></icons>
        <examples></examples>
        <download-section></download-section>
        <carousel></carousel>
    </div>
</template>
<script>
import Hero from "./components/Hero.vue";
import BasicElements from "./components/BasicElements.vue";
import Inputs from "./components/Inputs.vue";
import CustomControls from "./components/CustomControls.vue";
import Navigation from "./components/Navigation.vue";
import JavascriptComponents from "./components/JavascriptComponents.vue";
import Carousel from "./components/Carousel.vue";
import Icons from "./components/Icons.vue";
import Examples from "./components/Examples.vue";
import DownloadSection from "./components/DownloadSection.vue";

export default {
  name: "components",
  components: {
    Hero,
    BasicElements,
    Inputs,
    CustomControls,
    Navigation,
    JavascriptComponents,
    Carousel,
    Icons,
    Examples,
    DownloadSection
  }
};
</script>
