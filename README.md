# ibmq

## Usage

```shell
Usage:
  ibmq backend [flags]
  ibmq backend [command]

Available Commands:
  ls          A brief description of your command
  status      A brief description of your command

Flags:
  -h, --help   help for backend

Use "ibmq backend [command] --help" for more information about a command.
```

## Set up

▼ command

```shell
ibmq configure
```

▼ output

```shell
IBM Quantum API Token:xxxxxxxxxxx
API Token saved
```

## Backend List

▼ command

```
ibmq backend ls
```

▼ output

```shell
ls called
- ibm_oslo
- ibmq_qasm_simulator
- ibmq_manila
- simulator_extended_stabilizer
- simulator_statevector
- ibm_lagos
- ibmq_quito
- ibmq_belem
- ibm_perth
- ibmq_jakarta
- simulator_stabilizer
- ibm_nairobi
- ibmq_lima
- simulator_mps
```

### Detail Option

▼ command

```
ibmq backend ls -l
```

▼ output

```shell
ls called
- ibm_oslo
  Backend version:
  State: true
  Message: available
  Length Queue: 41

- ibmq_qasm_simulator
  Backend version:
  State: true
  Message: active
  Length Queue: 1

- ibmq_manila
  Backend version:
  State: true
  Message: available
  Length Queue: 128

- simulator_extended_stabilizer
  Backend version:
  State: true
  Message: active
  Length Queue: 1

- simulator_statevector
  Backend version:
  State: true
  Message: active
  Length Queue: 1

- ibm_lagos
  Backend version:
  State: true
  Message: available
  Length Queue: 515

- ibmq_quito
  Backend version:
  State: true
  Message: available
  Length Queue: 33

- ibmq_belem
  Backend version:
  State: true
  Message: available
  Length Queue: 20

- ibm_perth
  Backend version:
  State: true
  Message: available
  Length Queue: 80

- ibmq_jakarta
  Backend version:
  State: true
  Message: available
  Length Queue: 25

- simulator_stabilizer
  Backend version:
  State: true
  Message: active
  Length Queue: 1

- ibm_nairobi
  Backend version:
  State: true
  Message: available
  Length Queue: 60

- ibmq_lima
  Backend version:
  State: true
  Message: available
  Length Queue: 378

- simulator_mps
  Backend version:
  State: true
  Message: active
  Length Queue: 1
```

## Backend Status

▼ command

```shell
ibmq backend status -d ibmq_lima
```

▼ output

```shell
status called
- ibmq_lima
  Backend version:
  State: true
  Message: available
  Length Queue: 378
```
