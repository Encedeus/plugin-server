<script>

    export let pluginName
    export let style
    const pluginPage = `/plugin/${pluginName}/`
</script>

<a href={pluginPage} class="w-fit">
    <p class="w-fit" style={style} >{pluginName}</p>
</a>




