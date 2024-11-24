<template>
  <h3>{{ text }}</h3>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'AuthLabel',
  props: ['text'],
});
</script>
