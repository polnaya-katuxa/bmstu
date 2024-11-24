<template>
  <div class="modal fade" id="staticBackdrop4" tabindex="-1" aria-labelledby="exampleModalLabel4"
       aria-hidden="true">
    <div class="modal-dialog modal-lg modal-dialog-centered d-flex justify-content-center">
      <div class="modal-content w-75">
        <ModalHeader :header="'New post'" :id="opt+curLogin"/>

        <ModalBody :placeholder="'Your post content'" :text="'Post content'"
                   :submit="'Publish'" :action="'onSubmitPost'"/>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { mapActions, mapGetters } from 'vuex';
import ModalHeader from '@/components/ModalHeader.vue';
import ModalBody from '@/components/ModalBody.vue';

export default defineComponent({
  name: 'PublishModal',
  props: ['opt'],
  components: { ModalBody, ModalHeader },
  data() {
    return {
      perms: false,
      content: '',
    };
  },
  methods: {
    ...mapActions(['publishPost']),
    // onSubmit() {
    //   this.publishPost({ content: this.content, perms: this.perms });
    //
    //   const obj = document.getElementById(`close${this.curLogin}`);
    //   if (obj !== null) {
    //     obj.click();
    //   }
    //
    //   this.content = '';
    //   this.perms = false;
    // },
  },
  computed: {
    ...mapGetters(['curLogin']),
  },
});
</script>

<style>
.check-pink {
  border-color: #d7c1f1 !important;
  background-color: #fff !important;
}
.check-pink:checked {
  border-color: #d7c1f1 !important;
  background-color: #d7c1f1 !important;
}
.text-pink {
  min-height: 300px !important;
  border-color: #d7c1f1;
}
.header-pink {
  background-color: #d7c1f1;
}
</style>
