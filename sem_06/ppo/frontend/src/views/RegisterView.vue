<template>
  <div class="position-relative vh-100">

    <div class="row" style="margin-top: 380px">
      <div class="col-md-4">

      </div>

      <div class="col-md-4 px-5">
        <form class="position-absolute start-50 translate-middle w-25 needs-validation"
              @submit.prevent="onSubmit">

          <MainLogo/>

          <AuthLabel :text="'Registration'"/>

          <AuthField :placeholder="'Your login'" :text="'Login'"
                     :invalid="'Please choose a login.'"
                     v-model:inputValue="loginForm"/>
          <AuthField :placeholder="'Your password'" :text="'Password'"
                     :invalid="'Please choose a password.'"
                     v-model:inputValue="passwordForm"/>
          <AuthField :placeholder="'Your email'" :text="'Email'"
                     :invalid="'Please choose an email.'"
                     v-model:inputValue="emailForm"/>
          <AuthField :placeholder="'Your picture'" :text="'Picture'"
                     :invalid="'Please choose a picture.'"
                     v-model:inputValue="pictureForm"/>
          <AuthField :placeholder="'Your description'" :text="'Description'"
                     :invalid="'Please choose a description.'"
                     v-model:inputValue="descriptionForm"/>

          <div class="row w-100 me-0 ms-0">
            <AuthButton :text="'Submit'"/>
            <AuthLink :whereto="'/login'" :text="'Registered already?'"/>
          </div>
        </form>
      </div>

      <div class="col-md-4">

      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { mapActions } from 'vuex';
import AuthField from '@/components/AuthField.vue';
import MainLogo from '@/components/Logo.vue';
import AuthLabel from '@/components/AuthLabel.vue';
import AuthLink from '@/components/AuthLink.vue';
import AuthButton from '@/components/AuthButton.vue';

export default defineComponent({
  name: 'RegisterView',
  components: {
    AuthButton,
    AuthLink,
    AuthLabel,
    MainLogo,
    AuthField,
  },
  data() {
    return {
      loginForm: '',
      passwordForm: '',
      emailForm: '',
      pictureForm: '',
      descriptionForm: '',
    };
  },
  methods: {
    ...mapActions(['register']),
    onSubmit() {
      this.register({
        login: this.loginForm,
        password: this.passwordForm,
        email: this.emailForm,
        picture: this.pictureForm,
        description: this.descriptionForm,
      });
    },
  },
});
</script>

<style scoped>

</style>
