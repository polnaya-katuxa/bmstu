<template>
  <div class="ms-3" style="margin-top: 85px; color: #000;">
    <h5>{{ textBigger }}</h5>
    <p>{{ textSmaller }}</p>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'ProfileCardHeaderMainInfo',
  props: ['textBigger', 'textSmaller'],
});
</script>

<style scoped>

</style>
