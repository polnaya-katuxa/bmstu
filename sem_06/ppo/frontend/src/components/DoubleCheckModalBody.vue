<template>
  <div class="modal-body p-4">
    <!-- eslint-disable-next-line -->
    <label class="form-check-label" style="margin-bottom: 24px" for="flexSwitchCheckDefault"
    >{{ text }}</label>
    <div class="d-flex justify-content-between">
      <DoubleCheckButton @deleted="onDeleted" :text="'Delete'" :action="'onClickDelete'"
                         :id="id"/>
      <DoubleCheckButton :text="'Cancel'" :action="'onClickCancel'" :id="id"/>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import DoubleCheckButton from '@/components/DoubleCheckButton.vue';

export default defineComponent({
  name: 'DoubleCheckModalBody',
  components: { DoubleCheckButton },
  props: ['text', 'id'],
  methods: {
    onDeleted() {
      this.$emit('deleted');
    },
  },
});
</script>
