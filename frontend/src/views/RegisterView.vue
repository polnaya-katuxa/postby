<template>
  <div class="position-relative vh-100">

    <div class="row" style="margin-top: 380px">
      <div class="col-md-4">

      </div>

      <div class="col-md-4 px-5">
        <form class="position-absolute start-50 translate-middle w-25 needs-validation"
              @submit.prevent="onSubmit">

          <div class="d-flex justify-content-center">
            <img src="../assets/logo.png"
                 style="min-width: 160px; min-height: 160px; max-width: 160px;
                             max-height: 160px"
                 alt=""/>
          </div>

          <h3>Registration</h3>
          <div class="form-floating mb-3">
            <input type="text" pattern=".{4,20}" minlength="4" title="4 characters minimum"
                   class="form-control" id="inputLogin" required
                   placeholder="Your login" v-model="loginForm">
            <!-- eslint-disable-next-line -->
            <label for="inputLogin">
              Login
            </label>
          </div>

          <div class="form-floating mb-3">
            <input type="password" pattern=".{8,20}" minlength="8" required
                   title="8 characters minimum"
                   class="form-control" id="inputPassword" placeholder="Your password"
                   v-model="passwordForm">
            <!-- eslint-disable-next-line -->
            <label for="inputPassword">
              Password
            </label>
          </div>

          <div class="form-floating mb-3">
            <input type="email" required
                   class="form-control" id="inputEmail" placeholder="Your email"
                   v-model="emailForm">
            <!-- eslint-disable-next-line -->
            <label for="inputEmail">
              Email
            </label>
          </div>

          <div class="form-floating mb-3">
            <input type="url" required
                   class="form-control" id="inputPicture" placeholder="Your picture"
                   v-model="pictureForm">
            <!-- eslint-disable-next-line -->
            <label for="inputPicture">
              Picture
            </label>
          </div>

          <div class="form-floating mb-3">
            <input type="text" required class="form-control" pattern=".{0,150}"
                   id="inputDescription" placeholder="Your description"
                   v-model="descriptionForm">
            <!-- eslint-disable-next-line -->
            <label for="inputDescription">
              Description
            </label>
          </div>
          <div class="row w-100 me-0 ms-0">
            <button type="submit" class="btn btn-pink col">Submit</button>
            <router-link to="/login" class="col btn btn-pink-link">Registered already?</router-link>
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

export default defineComponent({
  name: 'RegisterView',
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
