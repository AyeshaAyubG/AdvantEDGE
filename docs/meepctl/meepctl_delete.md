## meepctl delete

Delete containers from the K8s cluster

### Synopsis

Delete containers from the K8s cluster

AdvantEDGE is composed of a collection of micro-services (a.k.a the groups).

Delete command removes a group of containers from the K8s cluster.

```
meepctl delete <group> [flags]
```

### Examples

```
  # Delete dependency containers
  meepctl delete dep
  # Delete only AdvantEDGE core containers
  meepctl delete core

Valid Targets:
  * dep
  * core
```

### Options

```
  -h, --help   help for delete
```

### Options inherited from parent commands

```
  -t, --time      Display timing information
  -v, --verbose   Display debug information
```

### SEE ALSO

* [meepctl](meepctl.md)	 - meepctl - CLI application to control the AdvantEDGE platform

###### Auto generated by spf13/cobra on 13-Dec-2022
