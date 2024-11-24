<template>
  <div class="flex-shrink-0">
    <img v-bind:src="pic"
         style="min-width: 50px; min-height: 50px; max-width: 50px;
                         max-height: 50px"
         class="align-self-start mr-3 rounded-circle"
         alt=""/>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'CardAvatar',
  props: ['pic'],
});
</script>

<style scoped>

</style>
