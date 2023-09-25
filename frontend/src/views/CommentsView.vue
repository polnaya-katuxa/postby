<template>
  <BaseSection>
    <PostCard :post="getPost"/>

    <div v-if="allComments.length">
      <div v-for="comm in allComments" :key="comm.id" class="card mb-4">
        <div class="card-body">
          <div class="d-flex mb-3">
            <div class="flex-shrink-0">
              <img v-bind:src="comm.commentator.picture"
                   style="min-width: 50px; min-height: 50px; max-width: 50px;
                           max-height: 50px"
                   class="align-self-start mr-3 rounded-circle"
                   alt=""/>
            </div>
            <div class="flex-grow-1 ms-3" style="padding-top: 0.125rem">
              <router-link style="color: black; text-decoration: none"
                           v-bind:to="'/profile/' + comm.commentator.login"
                            >{{ comm.commentator.login }}
              </router-link>
              <div class="text-muted small">{{ new Date(comm.pubTime).toLocaleString() }}
              </div>
            </div>

            <div>
              <button class="btn btn-sm btn-reaction" type="button" id="deleteCommButton"
                  @click.prevent="uncomment({postID: getPost.id, commID: comm.id})"
                  v-if="curLogin == comm.commentator.login || getPost.author.login == curLogin">
                Delete
              </button>
            </div>
          </div>

          <div class="styled-scrollbars" style="overflow: auto">
            <div v-html="comm.content"></div>
          </div>
        </div>
      </div>
    </div>

    <div v-else>
      <div class="card">
        <div class="card-body">
          No comments
        </div>
      </div>
    </div>
    <CommentModal/>
  </BaseSection>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import * as openapi from '@/openapi';
import BaseSection from '@/components/BaseSection.vue';
import { mapActions, mapGetters } from 'vuex';
import CommentModal from '@/components/CommentModal.vue';
import PostCard from '@/components/PostCard.vue';

export default defineComponent({
  name: 'CommentsView',
  components: { PostCard, CommentModal, BaseSection },
  created() {
    const rl = this.$route.params.postID.toString();

    this.getComments(rl);
  },
  methods: {
    ...mapActions(['getComments', 'comment', 'uncomment']),
  },
  computed: {
    ...mapGetters(['allComments', 'curLogin', 'getPost']),
  },
});
</script>

<style>
.card  {
  border-color: #d7c1f1 !important;
}
</style>
