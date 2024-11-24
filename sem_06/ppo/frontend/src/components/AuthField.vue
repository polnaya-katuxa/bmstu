<template>
  <div class="form-floating mb-3 has-validation">
    <input type="text" pattern=".{4,20}" required minlength="4" class="form-control"
           id="inputLogin" :placeholder="placeholder" :value="inputValue"
           @input="$emit('update:inputValue', $event.target.value)">
    <!-- eslint-disable-next-line -->
    <label for="inputLogin">
      {{ text }}
    </label>
    <div class="invalid-feedback">
      {{ invalid }}
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'AuthField',
  props: ['placeholder', 'text', 'invalid', 'inputValue'],
});
</script>
