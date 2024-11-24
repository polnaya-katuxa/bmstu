<template>
  <div class="modal-header header-pink">
    <h5 class="modal-title" id="exampleModalLabel5"> {{ header }}</h5>
    <button type="button" :id="`close`+id" class="btn-close" data-bs-dismiss="modal"
            aria-label="Close"></button>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'ModalHeader',
  props: ['header', 'id'],
});
</script>

<style scoped>

</style>
