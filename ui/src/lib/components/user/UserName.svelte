<script lang="ts">
    export let username: string;
    export let className: string;
    let userPage = `/user/${username}`;

</script>

<a href={userPage} class={className}>{username}</a>