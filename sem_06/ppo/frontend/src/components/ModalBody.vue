<template>
  <div class="modal-body p-4">
    <form @submit.prevent="getSubmitFunc">

      <ModalTextArea v-model:content="content" :placeholder="placeholder" :text="text"/>

      <div v-if="location === `${basePath}profile/` + curLogin">
        <br>
        <PermsSwitch v-model:perms="perms"/>
      </div>

      <br>

      <button type="submit" class="btn btn-reaction btn-block">{{ submit }}</button>
    </form>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { mapActions, mapGetters } from 'vuex';
import ModalTextArea from '@/components/ModalTextArea.vue';
import PermsSwitch from '@/components/PermsSwitch.vue';

export default defineComponent({
  name: 'ModalBody',
  components: { PermsSwitch, ModalTextArea },
  props: ['placeholder', 'text', 'submit', 'action'],
  data() {
    return {
      perms: false,
      content: '',
    };
  },
  methods: {
    ...mapActions(['comment', 'publishPost']),
    onSubmitComment() {
      this.comment({ postID: this.getPost.id, content: this.content });

      console.error(`closecomment${this.getPost.id}`);

      const obj = document.getElementById(`closecomment${this.getPost.id}`);
      if (obj !== null) {
        console.error(obj);
        obj.click();
      }

      this.content = '';
    },
    onSubmitPost() {
      this.publishPost({ content: this.content, perms: this.perms });

      console.error(`closepost${this.curLogin}`);

      const obj = document.getElementById(`closepost${this.curLogin}`);
      if (obj !== null) {
        console.error(obj);
        obj.click();
      }

      this.content = '';
      this.perms = false;
    },
    getSubmitFunc() {
      if (this.action === 'onSubmitPost') {
        this.onSubmitPost();
      }
      if (this.action === 'onSubmitComment') {
        this.onSubmitComment();
      }
    },
  },
  computed: {
    ...mapGetters(['getPost', 'curLogin']),
    location() {
      return window.location.pathname;
    },
    basePath() {
      return process.env.BASE_URL;
    },
  },
});
</script>
