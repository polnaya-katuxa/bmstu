<template>
  <div class="form-floating">
              <textarea class="form-control text-pink" :placeholder=placeholder
                        id="floatingTextarea2" style="height: 100px"
                        :value="content"
                        @input="$emit('update:content', $event.target.value)"></textarea>
    <!-- eslint-disable-next-line -->
    <label for="floatingTextarea2"> {{ text }} </label>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'ModalTextArea',
  props: ['placeholder', 'text', 'content'],
});
</script>

<style scoped>

</style>
