<template>
  <div class="form-check form-switch">
    <input class="form-check-input check-pink" type="checkbox" id="flexSwitchCheckDefault"
           @input="$emit('update:perms', $event.target.value === 'on')">
    <!-- eslint-disable-next-line -->
    <label class="form-check-label" for="flexSwitchCheckDefault"
    >Available after subscription</label>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'PermsSwitch',
  props: ['perms'],
});
</script>

<style scoped>

</style>
