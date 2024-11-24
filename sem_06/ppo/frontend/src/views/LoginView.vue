<template>
  <div class="position-relative vh-100">

    <div class="row" style="margin-top: 180px">
      <div class="col-md-4">

      </div>

      <div class="col-md-4 px-5">
        <form class="top-50 needs-validation"
              @submit.prevent="onSubmit">

          <MainLogo/>

          <AuthLabel :text="'Login'"/>

          <AuthField :placeholder="'Your login'" :text="'Login'" :invalid="'Please choose a login.'"
                     v-model:inputValue="loginForm"/>
          <AuthField :placeholder="'Your password'" :text="'Password'"
                     :invalid="'Please choose a password.'"
                     v-model:inputValue="passwordForm"/>

          <div class="row w-100 me-0 ms-0">
            <AuthButton :text="'Submit'"/>
            <AuthLink :whereto="'/register'" :text="'Not registered yet?'"/>
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
  name: 'LoginView',
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
    };
  },
  methods: {
    ...mapActions(['login']),
    onSubmit() {
      console.error(this.loginForm, this.passwordForm);
      this.login({
        login: this.loginForm,
        password: this.passwordForm,
      });
    },
  },
});
</script>
