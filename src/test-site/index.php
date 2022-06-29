<?php

header("X-Robots-Tag: noindex, nofollow", true);
echo 'Hi there, welcome to the Scopic Software test site! </br>';
echo "</br>Your IP address is: ";
echo $_SERVER['REMOTE_ADDR'];

echo "</br>Server IP address is: ";
echo $_SERVER['SERVER_ADDR'];

echo "</br>";

if ($_SERVER['SERVER_ADDR'] == "10.253.255.122"){
        echo "</br>Great! = ) You are successfully connected to the internal VPN network.";
} else {
        echo "</br>VPN status: disconnected";
}

?>