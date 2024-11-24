<template>
  <section class="gradient-custom-2" style="min-height: 100vh">
    <div class="container py-4 h-100">
      <div class="row d-flex justify-content-center h-100">
        <div class="col col-lg-9 col-xl-7">

          <NavBar/>

          <div class="container w-746px">
            <br>

            <slot></slot>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import NavBar from '@/components/NavBar.vue';

export default defineComponent({
  name: 'BaseSection',
  components: {
    NavBar,
  },
});
</script>

<style>
.gradient-custom-2 {
  /* fallback for old browsers */
  background: #fbc2eb;

  /* Chrome 10-25, Safari 5.1-6 */
  background: -webkit-linear-gradient(to right, rgba(251, 194, 235, 1), rgba(166, 193, 238, 1));

  /* W3C, IE 10+/ Edge, Firefox 16+, Chrome 26+, Opera 12+, Safari 7+ */
  background: linear-gradient(to right, rgba(251, 194, 235, 1), rgba(166, 193, 238, 1))
}
</style>
