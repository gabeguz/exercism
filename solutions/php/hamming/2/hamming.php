<?php
/**
 * Distance calculates the Hamming distance between two DNA strands.
 *
 * @param  string $a DNA strand (ex: GACTCCA)
 * @param  string $b DNA strand (ex: CACTCCG)
 * @throws InvalidArgumentException if the two arguments are not the same
 *	length
 * @return int The hamming distance between the two strings
 */
function distance($a, $b)
{
    if (strlen($a) !== strlen($b)) {
        throw(new InvalidArgumentException('DNA strands must be of equal length.'));
    }

    $distance = 0;
    for ($i = 0; $i < strlen($a); $i++) {
        if ($a[$i] !== $b[$i]) {
            $distance++;
        }
    }

    return $distance;
}