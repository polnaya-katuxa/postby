<template>
  <div class="modal fade" id="staticBackdrop5" tabindex="-1" aria-labelledby="exampleModalLabel5"
       aria-hidden="true">
    <div class="modal-dialog modal-lg modal-dialog-centered d-flex justify-content-center">
      <div class="modal-content w-75">
        <div class="modal-header header-pink">
          <h5 class="modal-title" id="exampleModalLabel5">New comment</h5>
          <button type="button" id="close" class="btn-close" data-bs-dismiss="modal"
                  aria-label="Close"></button>
        </div>
        <div class="modal-body p-4">
          <form @submit.prevent="onSubmit">

            <div class="form-floating">
              <textarea class="form-control text-pink" placeholder="Leave a comment here"
                        id="floatingTextarea2" style="height: 100px" v-model="content"></textarea>
              <!-- eslint-disable-next-line -->
              <label for="floatingTextarea2">Comment content</label>
            </div>

            <br>

            <button type="submit" class="btn btn-reaction btn-block">Comment</button>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { mapActions, mapGetters } from 'vuex';

export default defineComponent({
  name: 'CommentModal',
  data() {
    return {
      content: '',
    };
  },
  methods: {
    ...mapActions(['comment']),
    onSubmit() {
      this.comment({ postID: this.getPost.id, content: this.content });

      const obj = document.getElementById('close');
      if (obj !== null) {
        obj.click();
      }

      this.content = '';
    },
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
