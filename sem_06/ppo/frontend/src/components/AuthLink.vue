<template>
  <router-link :to=whereto class="col btn btn-pink-link">{{ text }}</router-link>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'AuthLink',
  props: ['whereto', 'text'],
});
</script>
