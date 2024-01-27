<script>

    export let pluginName
    export let style
    const pluginPage = `/plugin/${pluginName}/`
</script>

<a href={pluginPage}>
    <p class="pluginName" style={style} >{pluginName}</p>
</a>




