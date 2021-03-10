#
import paramiko
import getpass  #getpass是隐藏密码

def ssh_connect():
    host_ip = '127.0.0.1'
    user_name = 'root'
    host_port ='23'
    password='ljd620904'

    # 待执行的命令
    sed_command = "sed -i 's/123/abc/g' /root/test/test.txt"
    ls_command = "ls /root/test/"
    who_command="whoami"
    su_command="su - opc"

    # 注意：依次执行多条命令时，命令之间用分号隔开
    command = sed_command+";"+ls_command

    # SSH远程连接
    ssh = paramiko.SSHClient()   #创建sshclient
    ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())  #指定当对方主机没有本机公钥的情况时应该怎么办，AutoAddPolicy表示自动在对方主机保存下本机的秘钥
    ssh.connect(host_ip, host_port, user_name, password)
    ssh.
    # 执行命令并获取执行结果
    stdin, stdout, stderr = ssh.exec_command(who_command)
    print(stdout.read())
    stdin, stdout, stderr = ssh.exec_command(su_command)
    print(stdout.read())
    stdin, stdout, stderr = ssh.exec_command(who_command)
    print(stdout.read())
    out = stdout.readlines()
    err = stderr.readlines()
    
    ssh.close()

    return out,err



if __name__ == '__main__':
    result = ssh_connect()
    print(result)