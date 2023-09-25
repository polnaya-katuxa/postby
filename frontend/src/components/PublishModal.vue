<template>
  <div class="modal fade" id="staticBackdrop4" tabindex="-1" aria-labelledby="exampleModalLabel4"
       aria-hidden="true">
    <div class="modal-dialog modal-lg modal-dialog-centered d-flex justify-content-center">
      <div class="modal-content w-75">
        <div class="modal-header header-pink">
          <h5 class="modal-title" id="exampleModalLabel4">New post</h5>
          <button type="button" id="close" class="btn-close" data-bs-dismiss="modal"
                  aria-label="Close"></button>
        </div>
        <div class="modal-body p-4">
          <form @submit.prevent="onSubmit">

            <div class="form-floating">
              <textarea class="form-control text-pink" placeholder="Your post content"
                        id="floatingTextarea2" style="height: 100px" v-model="content"></textarea>
              <!-- eslint-disable-next-line -->
              <label for="floatingTextarea2">Post content</label>
            </div>

            <br>

            <div class="form-check form-switch">
              <input class="form-check-input check-pink" type="checkbox" id="flexSwitchCheckDefault"
                     v-model="perms">
              <!-- eslint-disable-next-line -->
              <label class="form-check-label" for="flexSwitchCheckDefault"
              >Available after subscription</label>
            </div>

            <br>

            <button type="submit" class="btn btn-reaction btn-block">Publish</button>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { mapActions } from 'vuex';

export default defineComponent({
  name: 'PublishModal',
  data() {
    return {
      perms: false,
      content: '',
    };
  },
  methods: {
    ...mapActions(['publishPost']),
    onSubmit() {
      this.publishPost({ content: this.content, perms: this.perms });

      const obj = document.getElementById('close');
      if (obj !== null) {
        obj.click();
      }

      this.content = '';
      this.perms = false;
    },
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
