<template>
  <div class="modal fade" :id="'modal-' + id" tabindex="-1" aria-labelledby="exampleModalLabel5"
       aria-hidden="true">
    <div class="modal-dialog modal-lg modal-dialog-centered d-flex justify-content-center">
      <div class="modal-content w-75">

        <ModalHeader :header="'Are you sure?'" :id="id"/>

        <DoubleCheckModalBody @deleted="onDeleted" :text="'You won\'t be able to undo this action'"
                              :id="id"/>

      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import ModalHeader from '@/components/ModalHeader.vue';
import DoubleCheckModalBody from '@/components/DoubleCheckModalBody.vue';

export default defineComponent({
  name: 'DoubleCheckModal',
  components: { DoubleCheckModalBody, ModalHeader },
  props: ['id'],
  methods: {
    onDeleted() {
      this.$emit('deleted');
    },
  },
});
</script>
