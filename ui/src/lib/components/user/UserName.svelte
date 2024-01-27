<script lang="ts">
    export let username: string;
    export let style: string;
    let userPage = `/user/${username}`;

</script>

<a href={userPage}>
    <p class="userName" style={style}>{username}</p>
</a>