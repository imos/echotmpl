# echotmpl
Generic template engine.

## Example

Input:
```php
<?
#include <string>
#include <vector>

std::string List(const std::vector<std::string>& v) {
    std::string output;
    auto echo = [&output](const std::string& s) { output += s; };
?>
<html><body>
<ul>
<?
    for (const std::string& s : v) {
?>
<li>Item <?=s?></li>
<?
    }
?>
</li>
</body></html>
<?
}
```

Output:
```cpp
#include <string>
#include <vector>

std::string List(const std::vector<std::string>& v) {
    std::string output;
    auto echo = [&output](const std::string& s) { output += s; };
echo("\x0a<html><body>\x0a<ul>\x0a");

    for (const std::string& s : v) {
echo("\x0a<li>Item ");
echo(s);
echo("</li>\x0a");

    }
echo("\x0a</li>\x0a</body></html>\x0a");

}
```

