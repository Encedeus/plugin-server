<script lang="ts">

    export let pluginName;
    export let className;
    const pluginPage = `/plugin/${pluginName}/`;
</script>

<a href={pluginPage} class={className}>{pluginName}</a>




