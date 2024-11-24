<template>
  <div class="modal fade" id="staticBackdrop5" tabindex="-1" aria-labelledby="exampleModalLabel5"
       aria-hidden="true">
    <div class="modal-dialog modal-lg modal-dialog-centered d-flex justify-content-center">
      <div class="modal-content w-75">
        <ModalHeader :header="'New comment'" :id="opt+getPost.id"/>

        <ModalBody :placeholder="'Leave a comment here'" :submit="'Comment'"
                   :text="'Comment content'" :action="'onSubmitComment'"/>
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
  name: 'CommentModal',
  props: ['opt'],
  components: { ModalBody, ModalHeader },
  data() {
    return {
      content: '',
    };
  },
  methods: {
    ...mapActions(['comment']),
    // onSubmit() {
    //   this.comment({ postID: this.getPost.id, content: this.content });
    //
    //   const obj = document.getElementById(`close${this.getPost.id}`);
    //   if (obj !== null) {
    //     obj.click();
    //   }
    //
    //   this.content = '';
    // },
  },
  computed: {
    ...mapGetters(['getPost']),
  },
});
</script>

<style>
.text-pink {
  min-height: 300px !important;
  border-color: #d7c1f1;
}
.header-pink {
  background-color: #d7c1f1;
}
</style>
