<template>
  <button type="submit" class="btn btn-reaction btn-block"
          @click.prevent="getClickFunc" data-bs-dismiss="modal">{{ text }}</button>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'DoubleCheckButton',
  props: ['text', 'action', 'id'],
  methods: {
    onClickDelete() {
      this.$emit('deleted');
      this.onClickCancel();
    },
    onClickCancel() {
      const obj = document.getElementById(`closecheck${this.id}`);
      if (obj !== null) {
        obj.click();
      }
    },
    getClickFunc() {
      if (this.action === 'onClickDelete') {
        this.onClickDelete();
      }
      if (this.action === 'onClickCancel') {
        this.onClickCancel();
      }
    },
  },
});
</script>
