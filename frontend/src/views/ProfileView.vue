<template>
  <BaseSection>
    <ProfileCard/>
  </BaseSection>

  <PublishModal/>

</template>

<script lang="ts">
// eslint-disable-next-line max-classes-per-file
import { defineComponent } from 'vue';
import BaseSection from '@/components/BaseSection.vue';
import { mapActions, mapGetters } from 'vuex';
import ProfileCard from '@/components/ProfileCard.vue';
import PublishModal from '@/components/PublishModal.vue';

export default defineComponent({
  name: 'ProfileView',
  components: { PublishModal, ProfileCard, BaseSection },
  data() {
    return {
      perms: false,
      content: '',
    };
  },
  mounted() {
    const rl = this.$route.params.login.toString();

    this.viewProfile(rl);
    this.getProfilePosts(rl);
  },
  created() {
    window.scrollTo(0, 0);
  },
  methods: {
    ...mapActions(['getProfilePosts', 'changePermsPost', 'deletePost', 'publishPost',
      'viewProfile', 'subscribe']),
    onSubmit() {
      this.publishPost({ content: this.content, perms: this.perms });

      const obj = document.getElementById('close');
      if (obj !== null) {
        obj.click();
      }
    },
    reactionClass(reacted: boolean): string {
      if (reacted) {
        return 'btn-my-reaction';
      }

      return 'btn-reaction';
    },
  },
  computed: {
    ...mapGetters(['allPosts', 'profile', 'subscribed', 'self']),
  },
});
</script>
