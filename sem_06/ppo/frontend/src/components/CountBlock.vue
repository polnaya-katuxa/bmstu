<template>
  <div>
    <p class="mb-1 h5">{{ count }}</p>
    <p class="small text-muted mb-0">{{ label }}</p>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'CountBlock',
  props: ['count', 'label'],
});
</script>
