<?php
    define("TEMPLATE", __DIR__ . "/../templates/README.md.php");
    define("OUTPUT_FILE", __DIR__ . "/../README.md");

    ob_start();
    $res_content = require_once TEMPLATE;
    $content = ob_get_clean();
    $content = file_put_contents(OUTPUT_FILE, $content);
    echo $res_content !== false && false !== $res ? "OK" : "ERROR";