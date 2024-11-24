<template>
  <div class="p-4 text-black" style="background-color: rgba(248, 242, 255, 0.99);">
    <div class="d-flex justify-content-end gap-3 text-center py-1">
      <CountBlock :count="stat1" :label="'Posts'"/>
      <CountBlock :count="stat2" :label="'Coins'"/>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import CountBlock from '@/components/CountBlock.vue';

export default defineComponent({
  name: 'ProfileStatsBlock',
  components: { CountBlock },
  props: ['stat1', 'stat2'],
});
</script>

<style scoped>

</style>
